package domain

import (
	"encoding/json"
)

type Funcionario struct {
	Id                int     `json:"id" gorm:"primaryKey;column:id"`
	CodUJ             uint    `json:"cod_uj" gorm:"default:1000200;column:cd_uj"`
	CodUPAG           string  `json:"cod_upag" gorm:"default:'MPC/PA;column:cd_upag"`
	Nome              string  `json:"nome" gorm:"column:nm_funcionario"`
	CPF               string  `json:"cpf" gorm:"column:cd_cpf"`
	Matricula         string  `json:"matricula" gorm:"column:cd_matricula"`
	Regime            uint    `json:"regime" gorm:"column:id_regime"`
	Cargo             string  `json:"cargo"  gorm:"column:ds_cargo"`
	NaturezaCargo     uint    `json:"natureza_cargo" gorm:"column:id_natureza_cargo"`
	DataExercicio     string  `json:"data_exercicio" gorm:"column:dt_exercicio"`
	DataAposentadoria string  `json:"data_aposentadoria" gorm:"column:dt_aposentadoria"`
	DataExclusao      string  `json:"data_exclusao" gorm:"column:dt_exclusao"`
	Jornada           uint    `json:"jornada" gorm:"column:nr_jornada"`
	CategoriaSituacao uint    `json:"categoria_situacao" gorm:"column:cd_categoria_situacao"`
	MesFolhaAno       string  `json:"mes_folha_ano" gorm:"default:10/2022;column:mes_folha_ano"`
	ValorBruto        float32 `json:"valor_bruto" gorm:"column:vl_bruto"`
}


func (f *Funcionario) ToJson() (string, error) {
	b, err := json.Marshal(f)
	if err != nil {
		return "",err
	}
	return string(b), nil
}

func FuncionariosToJson(fs []Funcionario) (string, error) {
	b, err := json.Marshal(fs)
	if err != nil {
		return "",err
	}
	return string(b), nil
}
