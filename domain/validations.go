package domain

func (tm *taskManager) validateUserID(id int) bool {
	if _, err := tm.userRepo.FindWithID(id); err != nil {
		return false
	}

	return true
}

func (tm *taskManager) validateProjectOwner(userID, projectID int) bool {
	proj, err := tm.projectRepo.FindWithID(projectID)
	if err != nil {
		return false
	}

	if proj.UserID != userID {
		return false
	}

	return true
}

func (tm *taskManager) validateTaskOwner(userID, taskID int) bool {
	tsk, err := tm.taskRepo.FindWithID(taskID)
	if err != nil {
		return false
	}

	return tm.validateProjectOwner(userID, tsk.ProjectID)
}
