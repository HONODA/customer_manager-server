/*test"
CALL sp_viewPage(
'*'#查询字段
,'userupdatelog'#表名
,'1=1'#条件
,'Id desc'#排序
,1 #页码
,20 #每页记录数
,@totalcount #输出总记录数
,@pagecount #输出用页数
);
SELECT @totalcount,@pagecount;
*/

DROP PROCEDURE IF EXISTS `sp_viewPage`;

CREATE PROCEDURE sp_viewPage(
_fields VARCHAR(1000), #要查询的字段，用逗号(,)分隔
_tables TEXT, #要查询的表
_where VARCHAR(2000), #查询条件
_orderby VARCHAR(200), #排序规则
_pageindex INT, #查询页码
_pageSize INT, #每页记录数
/*_sumfields VARCHAR(200),#求和字段 */
#输出参数
OUT _totalcount INT, #总记录数
OUT _pagecount INT #总页数
/* OUT _sumResult VARCHAR(2000)#求和结果 */
)
BEGIN
#140529-xxj-分页存储过程
#计算起始行号
SET @startRow = _pageSize * (_pageIndex - 1);
SET @pageSize = _pageSize;
SET @rowindex = 0; #行号

#合并字符串
SET @strsql = CONCAT(
#'select sql_calc_found_rows @rowindex:=@rowindex+1 as rownumber,' #记录行号
'select sql_calc_found_rows '
,_fields
,' from '
,_tables
,CASE IFNULL(_where, '') WHEN '' THEN '' ELSE CONCAT(' where ', _where) END
,CASE IFNULL(_orderby, '') WHEN '' THEN '' ELSE CONCAT(' order by ', _orderby) END
,' limit '
,@startRow
,','
,@pageSize
);

PREPARE strsql FROM @strsql;#定义预处理语句
EXECUTE strsql; #执行预处理语句
DEALLOCATE PREPARE strsql; #删除定义
#通过 sql_calc_found_rows 记录没有使用 limit 语句的记录，使用 found_rows() 获取行数
SET _totalcount = FOUND_ROWS();

#计算总页数
IF (_totalcount <= _pageSize) THEN
SET _pagecount = 1;
ELSE IF (_totalcount % _pageSize > 0) THEN
SET _pagecount = _totalcount DIV _pageSize + 1;
ELSE
SET _pagecount = _totalcount DIV _pageSize;
END IF;
END IF;
SELECT _totalcount,_pagecount;
END