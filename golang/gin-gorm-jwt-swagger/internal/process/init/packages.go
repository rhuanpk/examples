package init

import (
	// Imports of all packages project.
	_ "backend/internal/methods/auth"
	_ "backend/internal/methods/user"
	_ "backend/internal/models"
	_ "backend/internal/operations/database"
	_ "backend/pkg/docs"
)
