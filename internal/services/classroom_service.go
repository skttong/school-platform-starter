package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"school/internal/models"
)

// Dummy classroom service with pgx.ErrNoRows fixes
