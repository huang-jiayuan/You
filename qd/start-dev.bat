@echo off
echo Starting frontend development server...
echo.
echo Frontend will run on: http://localhost:3000
echo API requests will be proxied to:
echo   - Room service: http://localhost:8080
echo   - User service: http://localhost:8081
echo.
echo Make sure both backend services are running:
echo   - Room service on port 8080
echo   - User service on port 8081
echo.

npm run dev

pause