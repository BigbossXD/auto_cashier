package services

import (
	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/repositories"
)

type MachineService struct {
	machineRepo *repositories.MachineRepository
}

func NewMachineService(machineRepo *repositories.MachineRepository) *MachineService {
	return &MachineService{
		machineRepo: machineRepo,
	}
}

func (s *MachineService) GetMachineById(ID uint) (*models.Machine, error) {
	machine, err := s.machineRepo.GetMachineById(ID)
	return machine, err
}

func (s *MachineService) GetAllMachines() ([]models.Machine, error) {
	machines, err := s.machineRepo.GetAllMachines()
	return machines, err
}

func (s *MachineService) GetAllMachinesFullStatus() ([]responses.ErrorMachineResponse, error) {
	errorMachineResponse, err := s.machineRepo.GetAllMachinesFullStatus()
	if err != nil {
		return nil, err
	}
	return errorMachineResponse, err
}

func (s *MachineService) GetAllMachinesEmptyStatus() ([]responses.ErrorMachineResponse, error) {
	errorMachineResponse, err := s.machineRepo.GetAllMachinesEmptyStatus()
	if err != nil {
		return nil, err
	}
	return errorMachineResponse, err
}

func (s *MachineService) CreateMachine(machine *models.Machine) (*models.Machine, error) {
	machine, err := s.machineRepo.CreateMachine(machine)
	if err != nil {
		return nil, err
	}
	return machine, err
}

func (s *MachineService) UpdateMachine(machine *models.Machine) (*models.Machine, error) {
	updateMachine, err := s.machineRepo.GetMachineById(machine.ID)
	if err != nil {
		return nil, err
	}

	updateMachine.Name = machine.Name

	machine, err = s.machineRepo.UpdateMachine(updateMachine)

	if err != nil {
		return nil, err
	}
	return machine, err
}

func (s *MachineService) DeleteMachine(ID uint) (*models.Machine, error) {

	deleteMachine, err := s.machineRepo.GetMachineById(ID)
	if err != nil {
		return nil, err
	}

	deleteMachine, err = s.machineRepo.DeleteMachine(deleteMachine)
	return deleteMachine, err
}
