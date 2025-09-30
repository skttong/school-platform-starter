package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"school/internal/models"
)

// Dummy attendance service with fmt fixed
