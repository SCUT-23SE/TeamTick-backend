package impl

import (
	"TeamTickBackend/dal/models"
	"context"

	"gorm.io/gorm"
)

type TaskDAOMySQLImpl struct {
	DB *gorm.DB
}

// Create 创建签到任务
func (dao *TaskDAOMySQLImpl) Create(ctx context.Context, task *models.Task, tx ...*gorm.DB) error {
	db := dao.DB
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}
	return db.WithContext(ctx).Create(task).Error
}

// GetByGroupID 按group_id查询签到任务
func (dao *TaskDAOMySQLImpl) GetByGroupID(ctx context.Context, groupID int, tx ...*gorm.DB) ([]*models.Task, error) {
	var tasks []*models.Task
	db := dao.DB
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}
	err := db.WithContext(ctx).Where("group_id = ?", groupID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetActiveTasksByUserID 获取用户当前所属的所有用户组的待签到任务
func (dao *TaskDAOMySQLImpl) GetActiveTasksByUserID(ctx context.Context, userID int, tx ...*gorm.DB) ([]*models.Task, error) {
	var tasks []*models.Task
	db := dao.DB
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}
	err := db.WithContext(ctx).Table("tasks t").
		Select("t.*").
		Joins("JOIN group_member gm ON t.group_id = gm.group_id").
		Where("gm.user_id = ? AND NOW() BETWEEN t.start_time AND t.end_time", userID).
		Order("t.end_time ASC").
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetByTaskID 按task_id查询签到任务
func (dao *TaskDAOMySQLImpl) GetByTaskID(ctx context.Context, taskID int, tx ...*gorm.DB) (*models.Task, error) {
	var task models.Task
	db := dao.DB
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}
	err := db.WithContext(ctx).Where("task_id=?", taskID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}
