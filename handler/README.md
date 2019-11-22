### 请求处理器目录说明

- 默认文件加 `_handler.go` 作为后缀；
- 默认根据业务模块组织handler文件，如user模块:`user_handler.go` ;
- 可使用ginger-cli 的handler命令生成处理函数，如：`ginger-cli handler -f user -F SignIn -F SignUp`,该工具会自动生成处理器代码； 