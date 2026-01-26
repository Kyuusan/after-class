package task
import "tasklybe/pkg/db"

func GetTasks (page int, limit int)(*[]Task, int64, error){
	if page < 1 {
		page = 1 
	}

	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	var totalData int64 
	if err := db.DB.Model(&Task{}).Count(&totalData).Error; err != nil {
		return nil, 0, err
	} 

	offset := (page - 1) * limit
	var tasks []Task 
	if err := db.DB.Order("Created_at desc").Limit(limit).Offset(offset).Find(&tasks).Error; 
	err != nil {
		return nil, 0, err

	}

	return &tasks, totalData, nil
}

func GetDetailTask (id string) (*Task, error){
	var task Task 
	if err := db.DB.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &task, nil, 
}