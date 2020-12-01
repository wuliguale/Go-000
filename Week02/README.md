学习笔记

问题：
dao层中遇到sql.ErrNoRows错误时是否wrap后抛给上层？

答：
应该wrap后抛给上层，因为相比返回nil，这样没有歧义，而且便于上层判断
