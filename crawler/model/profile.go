package model

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Name         string //姓名
	Gender       string //性别
	Age          int    //年龄
	Height       int    //身高
	Weight       int    //体重
	Income       string //收入
	Marriage     string //婚姻状况
	Education    string //职业
	Occupation   string //教育
	Hokou        string //户口
	Xinzuo       string //星座
	House        string //房
	Car          string //车
	UserImageUrl string //用户照片
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}

func (p Profile) String() string {
	return fmt.Sprintf("name : %s ,userImageUrl : %s ,Gender : %s ,age : %d ,height : %d ,weight : %d ,"+
		"income : %s , marriage : %s , education : %s ,occupation : %s,hokou : %s , xinzuo : %s ,huose : %s ,car : %s",
		p.Name , p.UserImageUrl, p.Gender, p.Age, p.Height, p.Weight, p.Income, p.Marriage,
		p.Education, p.Occupation, p.Hokou, p.Xinzuo, p.House, p.Car)
}
