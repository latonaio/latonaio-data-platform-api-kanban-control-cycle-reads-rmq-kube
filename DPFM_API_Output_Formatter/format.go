package dpfm_api_output_formatter

import (
	"data-platform-api-kanban-control-cycle-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.KanbanControlCycle,
			&pm.KanbanControlCycleStrategy,
			&pm.PullRefillInterval,
			&pm.PullRefillIntervalUnit,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			KanbanControlCycle:         data.KanbanControlCycle,
			KanbanControlCycleStrategy: data.KanbanControlCycleStrategy,
			PullRefillInterval:         data.PullRefillInterval,
			PullRefillIntervalUnit:     data.PullRefillIntervalUnit,
			CreationDate:               data.CreationDate,
			LastChangeDate:             data.LastChangeDate,
			IsMarkedForDeletion:        data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}
