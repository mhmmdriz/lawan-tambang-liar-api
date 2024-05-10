package report_solution_process

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
)

type ReportSolutionProcessUseCase struct {
	repository entities.ReportSolutionProcessRepositoryInterface
	ai_api     entities.AIReportSolutionAPIInterface
}

func NewReportSolutionProcessUseCase(repository entities.ReportSolutionProcessRepositoryInterface, ai_api entities.AIReportSolutionAPIInterface) *ReportSolutionProcessUseCase {
	return &ReportSolutionProcessUseCase{
		repository: repository,
		ai_api:     ai_api,
	}
}

func (u *ReportSolutionProcessUseCase) Create(reportSolutionProcess *entities.ReportSolutionProcess) (entities.ReportSolutionProcess, error) {
	if reportSolutionProcess.ReportID == 0 || reportSolutionProcess.AdminID == 0 || reportSolutionProcess.Message == "" || reportSolutionProcess.Status == "" {
		return entities.ReportSolutionProcess{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Create(reportSolutionProcess)

	if err != nil {
		return entities.ReportSolutionProcess{}, constants.ErrInternalServerError
	}

	return *reportSolutionProcess, nil
}

func (u *ReportSolutionProcessUseCase) GetByReportID(reportID int) ([]entities.ReportSolutionProcess, error) {
	reportSolutionProcesses, err := u.repository.GetByReportID(reportID)

	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return reportSolutionProcesses, nil
}

func (u *ReportSolutionProcessUseCase) Delete(reportID int, reportSolutionProcessStatus string) (entities.ReportSolutionProcess, error) {
	reportSolutionProcess, err := u.repository.Delete(reportID, reportSolutionProcessStatus)

	if err != nil {
		return entities.ReportSolutionProcess{}, err
	}

	return reportSolutionProcess, nil
}

func (u *ReportSolutionProcessUseCase) Update(reportSolutionProcess entities.ReportSolutionProcess) (entities.ReportSolutionProcess, error) {
	if reportSolutionProcess.Message == "" {
		return entities.ReportSolutionProcess{}, constants.ErrAllFieldsMustBeFilled
	}

	reportSolution, err := u.repository.Update(reportSolutionProcess)

	if err != nil {
		return entities.ReportSolutionProcess{}, err
	}

	return reportSolution, nil
}

func (u *ReportSolutionProcessUseCase) GetMessageRecommendation(action string) (string, error) {
	// Membuat payload dari data pesan
	var messages []map[string]string
	if action == "verify" {
		messages = []map[string]string{
			{"role": "assistant", "content": "Anda sebagai admin website yang bertugas untuk memverifikasi sebuah laporan tambang liar, memberikan progress penyelesaian, dan menyelesaikan laporan"},
			{"role": "user", "content": "Saya seorang admin, ketika saya memverifikasi laporan, saya harus menambahkan sebuah pesan. Berikan saya contoh pesan yang baik dalam memverifikasi laporan tersebut ! Sebagai contoh : Laporan anda sudah diverifikasi, laporan anda akan segera diproses."},
		}
	} else if action == "progress" {
		messages = []map[string]string{
			{"role": "assistant", "content": "Anda sebagai admin website yang bertugas untuk memverifikasi sebuah laporan tambang liar, memberikan progress penyelesaian, dan menyelesaikan laporan"},
			{"role": "user", "content": "Saya seorang admin, ketika saya menambah sebuah progress, saya harus menambahkan sebuah pesan. Berikan saya contoh pesan yang baik saat menambahkan progress penyelesaian tersebut ! Sebagai contoh : Laporan anda sedang diproses, dst"},
		}
	} else if action == "finish" {
		messages = []map[string]string{
			{"role": "assistant", "content": "Anda sebagai admin website yang bertugas untuk memverifikasi sebuah laporan tambang liar, memberikan progress penyelesaian, dan menyelesaikan laporan"},
			{"role": "user", "content": "Saya seorang admin, ketika saya menyelesaikan laporan, saya harus menambahkan sebuah pesan. Berikan saya contoh pesan yang baik saat menyelesaikan laporan tersebut ! Sebagai contoh : Laporan anda telah selesai diproses, terima kasih atas kerjasamanya"},
		}
	} else {
		return "", constants.ErrActionNotFound
	}
	content, err := u.ai_api.GetChatCompletion(messages)

	if err != nil {
		return "", constants.ErrInternalServerError
	}

	return content, nil
}
