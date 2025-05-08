package sentio

// func SimulateLivePlugsWithSentio(livePlugs *signature.LivePlugs, apiKey string) (*models.Run, error) {
// 	client := NewSentioClient(apiKey)

// 	routerAddress := livePlugs.GetRouterAddress()

// 	var callData []byte
// 	var err error

// 	if livePlugs.Data != "" {
// 		callDataStr, err := hexutil.Decode(livePlugs.Data)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to decode call data: %v", err)
// 		}
// 		callData = callDataStr
// 	} else {
// 		callData, err = livePlugs.GetCallData()
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get call data: %v", err)
// 		}
// 	}

// 	fromAddress := common.HexToAddress(livePlugs.From)
// 	simulationID, err := client.SimulateTransaction(
// 		livePlugs.ChainId,
// 		fromAddress,
// 		routerAddress,
// 		callData,
// 		nil,
// 	)

// 	if err != nil {
// 		return nil, fmt.Errorf("simulation failed: %v", err)
// 	}

// 	callTrace, err := client.GetCallTrace(livePlugs.ChainId, simulationID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get call trace: %v", err)
// 	}

// 	// Get state diff but ignore it for now as it's not being used
// 	_, err = client.GetStateDiff(livePlugs.ChainId, simulationID)
// 	if err != nil {
// 		fmt.Printf("Warning: failed to get state diff: %v\n", err)
// 	}

// 	var errorReason *string
// 	if callTrace.Error != "" {
// 		reason := ExtractErrorFromCallTrace(callTrace)
// 		if reason != "" {
// 			errorReason = &reason
// 		}
// 	}

// 	// var stateChanges []models.StateChangeSummary
// 	// if stateDiff != nil {
// 	// 	summaries := SummarizeStateChanges(stateDiff)
// 	// 	for _, summary := range summaries {
// 	// 		stateChanges = append(stateChanges, models.StateChangeSummary{
// 	// 			Address:     summary.Address,
// 	// 			Description: summary.Description,
// 	// 			Type:        summary.Type,
// 	// 			ValueChange: summary.ValueChange,
// 	// 		})
// 	// 	}
// 	// }

// 	status := "success"
// 	var errString *string
// 	if !callTrace.Success || callTrace.Error != "" {
// 		status = "failed"
// 		errString = &callTrace.Error
// 	}

// 	var gasUsed uint64
// 	if callTrace.GasUsed > 0 {
// 		gasUsed = uint64(callTrace.GasUsed)
// 	}

// 	// callTraceJson, _ := json.Marshal(callTrace)
// 	// stateDiffJson, _ := json.Marshal(stateDiff)

// 	var errors []string
// 	if errString != nil {
// 		errors = []string{*errString}
// 	}

// 	run := &models.Run{
// 		LivePlugsId: livePlugs.Id,
// 		IntentId:    livePlugs.IntentId,
// 		From:        livePlugs.From,
// 		To:          routerAddress.Hex(),
// 		Status:      status,
// 		Error:       errorReason,
// 		Errors:      errors,
// 		GasUsed:     gasUsed,
// 		Data: models.RunOutputData{
// 			Raw: hexutil.MustDecode(callTrace.Output),
// 		},
// 	}

// 	return run, nil
// }
