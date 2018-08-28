package main

import (
	"fmt"
	"encoding/xml"
	"html"
	"strings"
)

type ResponseData struct {
	Code   int    `xml:"ResultCode"`
	Msg    string `xml:"ResultMessage"`
	Body   string `xml:"ResultList"`
}
type StuData struct {
	StudentEntity []StudentOne `xml:"ResultList>StudentInfo"`
}

type StudentOne struct {
	ID     int64   	   `xml:"StudentId"      json:"id"`               // 用户ID
	StudentName string `xml:"StudentName"    json:"real_name"` 	 // 用户姓名
	AreaName string    `xml:"AreaName"       json:"city"` 		 // 所属地区
	SchoolId int64     `xml:"SchoolId"       json:"school_id"` 	 // 学校ID
	SchoolName string  `xml:"SchoolName"     json:"name"` 		 // 学校名称
	GradeId int64      `xml:"GradeId"        json:"grade_id"` 	 // 年级id
	ClassId int64      `xml:"ClassId"        json:"unit_id"` 	 // 班级id
	ClassName string   `xml:"ClassName"      json:"class_name"`       // 班级名称
	StuSequence string `xml:"StuSequence"    json:"stu_sequence"`     // 学生编号
	FamilyInfo []struct{
		FamilyId    string  `xml:"FamilyId"   json:"family_id"`  // 家长ID
		FamilyName  string  `xml:"FamilyName" json:"family_name"`    // 家长姓名
		AccountId   string  `xml:"AccountId"  json:"account_id"`     // 账号ID  未生成账号的家长返回-1, 已生成账号的家长返回正常ID
	}`xml:"SubResultList>FamilyInfo"`
}

func main(){
	a := `
    <response>
  <ResultCode>0</ResultCode>
  <ResultMessage>成功</ResultMessage>
  <ResultList>
    <StudentInfo>
      <StudentId>392981</StudentId>
      <StudentName>张欣蕊</StudentName>
      <AreaName>南平</AreaName>
      <SchoolId>5796</SchoolId>
      <SchoolName>浦城县光明中心小学</SchoolName>
      <GradeId>3</GradeId>
      <GradeName>小学三年级</GradeName>
      <ClassId>257151</ClassId>
      <ClassName>小学三年级2班</ClassName>
      <StuSequence>159998749975</StuSequence>
      <SubResultList>
        <FamilyInfo>
          <FamilyId>425894</FamilyId>
          <FamilyName>张欣蕊家长</FamilyName>
          <AcountId>9172821</AcountId>
        </FamilyInfo>
      </SubResultList>
    </StudentInfo>
    <StudentInfo>
      <StudentId>392981</StudentId>
      <StudentName>张欣蕊1</StudentName>
      <AreaName>南平</AreaName>
      <SchoolId>5796</SchoolId>
      <SchoolName>浦城县光明中心小学</SchoolName>
      <GradeId>3</GradeId>
      <GradeName>小学三年级</GradeName>
      <ClassId>257151</ClassId>
      <ClassName>小学三年级2班</ClassName>
      <StuSequence>159998749975</StuSequence>
      <SubResultList>
        <FamilyInfo>
          <FamilyId>425894</FamilyId>
          <FamilyName>张欣蕊家长</FamilyName>
          <AcountId>9172821</AcountId>
        </FamilyInfo>
      </SubResultList>
    </StudentInfo>
  </ResultList>
</response>
`
	v := &StuData{}
	oAuthResponse := &ResponseData{}
	xml.NewDecoder(strings.NewReader(a)).Decode(oAuthResponse)
	fmt.Println(html.UnescapeString(a))
	xml.Unmarshal([]byte(html.UnescapeString(a)), v)
	fmt.Println(v)
	c := &v.StudentEntity
	fmt.Println(*c)
	
}
