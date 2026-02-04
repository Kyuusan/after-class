package task

import (
	"tasklybe/pkg/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

	return &task, nil
}

func CreateTask(input CreateTaskRequest) (*Task, error) {
	task := Task{
		ID:     uuid.NewString(),
		UserID: input.UserID,
		Title:  input.Title,
		Description:   input.Desc,
		Label:  input.Label,
	}

	if err := db.DB.Create(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func EditTask(id string, input EditTaskRequest) (*Task, error) {

	/**
	Kita inisialisasi variable task dengan entity Task, agar ORM tau kita sedang
	berinteraksi dengan table tasks
	**/
	var task Task
	
	// Cari datanya berdasarkan id apakah ada
	if err := db.DB.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	
	// Set value berdasarkan inputan user
	task.Title = input.Title
	task.Description = input.Desc
	task.Label = input.Label

	// Simpan ke database
	if err := db.DB.Save(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func DeleteTask(id string) error {
	
	/**
	Kita bisa langsung hapus data dengan method Delete() disertai entity
	yang ingin dihapus. tx = transaksi database
	**/
	tx := db.DB.Delete(&Task{}, "id = ?", id)

	// cek apakah ada transaksi db yang error.
	if tx.Error != nil {
		return tx.Error
	}
	
	// Check jika tidak ada sama sekali data yang terhapus, maka return notfound.
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// Kita tidak perlu mengembalikan nilai apapun
	return nil
}