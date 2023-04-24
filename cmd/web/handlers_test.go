package main 

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DerrickKirimi/Snippets/internal/assert"
)

func TestPing