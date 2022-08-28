package service

//这里就是我们的业务类、接口等相关信息存放
type IUserService interface {
    GetName(userId int) string
}

type UserService struct {
}

func (s UserService) GetName(uid int) string {
    if uid == 10{
        return "wangqiang"
    }
    return "xiaoqiang"
}