package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Student struct {
	name            string
	id              int
	list_of_courses []Course
}

func (student Student) viewCourses() {
	fmt.Println(student.list_of_courses)
}

func (student *Student) addCourse(course Course) {
	student.list_of_courses = append(student.list_of_courses, course)
}

func (student *Student) deleteCourse(course Course) {
	indexRemove := 2
	student.list_of_courses = remove(student.list_of_courses, indexRemove)
}

type Course struct {
	course_code string
	grade       int
}

func remove(slice []Course, s int) []Course {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	ExampleClient()
}

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	student := Student{name: "Bekzhan", id: 200107085}
	student.addCourse(Course{course_code: "CSS340", grade: 85})
	student.addCourse(Course{course_code: "INF423", grade: 90})
	student.addCourse(Course{course_code: "CSS440", grade: 92})
	course := Course{course_code: "INF324", grade: 78}
	student.deleteCourse(course)

	for _, course := range student.list_of_courses {
		err := rdb.HSet(ctx, student.name, course.course_code, course.grade).Err()
		if err != nil {
			fmt.Println(err)
		}
	}

	data, err := rdb.HGetAll(ctx, student.name).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(student.name, data)
}
