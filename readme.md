# Simple Wedding Management

A modern web application for wedding planning and management, built with Go and HTMX. This application helps couples plan their wedding by managing budgets, guest lists, wedding rundowns, and finding wedding organizers in the JABODETABEK area.

## Features

- **Budget Management**

  - Track wedding expenses
  - Set budget limits
  - Categorize expenses
  - Generate budget reports

- **Guest List Management**

  - Add and manage guest information
  - Track RSVP
  - Generate guest lists
  - Export guest data

- **Wedding Rundown**

  - Create and manage wedding timeline
  - Schedule activities
  - Track task completion
  - Set reminders

- **Wedding Organizer Directory**
  - Browse JABODETABEK wedding organizers
  - View organizer profiles
  - Contact information
  - Reviews and ratings

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTMX
- **Database**: Postgres
- **Template Engine**: tmpl

## Prerequisites

- Go 1.21 or higher
- PostgreSQL

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/simple-wedding-management.git
cd simple-wedding-management
```

2. Install dependencies:

```bash
go mod download
```

3. Set up environment variables:

```bash
cp config/example.config.yaml config/config.yaml
```

4. Run the application:

```bash
go run main.go
```

The application will be available at `http://localhost:8080`

## Project Structure

```
simple-wedding-management/
├── cmd/
│   └── main.go
├── config/
│   ├── databases/
├── internal/
│   ├── helpers/
│   └── models/
├── modules/
│   ├── handlers/
│   ├── models/
│   ├── services/
│   └── repository/
├── templates/
├── static/
├── go.mod
├── go.sum
└── README.md
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Go team for the amazing language
- HTMX team for the modern web approach
- All contributors who help improve this project
