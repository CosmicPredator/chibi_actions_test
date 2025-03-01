package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

var progress string
var updateStatus string
var notes string
var scoreString string

func handleUpdate(cmd *cobra.Command, args []string) {
	if len(args) == 2 {
		progress = args[1]
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(
			ui.ErrorText(errors.New("invalid media id. please provide a valid one")),
		)
	}

	var scoreFloat *float64
	if scoreString != "" {
		rawScoreFloat, err := strconv.ParseFloat(scoreString, 32)
		if err != nil {
			fmt.Println(ui.ErrorText(err))
			return
		}
		scoreFloat = &rawScoreFloat
	}

	params := viewmodel.MediaUpdateParams{
		IsNewAddition: false,
		MediaId:       id,
		Progress:      progress,
		Status:        updateStatus,
		StartDate:     "none",
	}
	// TODO: add a way to differentiate between an
	// empty notes value vs. an unset notes value
	if notes != "\n" {
		params.Notes = notes
	}
	if scoreFloat != nil {
		params.Score = float32(*scoreFloat)
	}
	err = viewmodel.HandleMediaUpdate(params)

	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var mediaUpdateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a list entry",
	Args:  cobra.MinimumNArgs(1),
	Run:   handleUpdate,
}

func init() {
	mediaUpdateCmd.Flags().StringVarP(
		&progress,
		"progress",
		"p",
		"",
		"The number of episodes/chapter to update",
	)
	mediaUpdateCmd.Flags().StringVarP(
		&updateStatus, "status", "s", "none", "Status of the media. Can be 'watching/w or reading/r', 'planning/p', 'completed/c', 'dropped/d', 'paused/ps'",
	)
	mediaUpdateCmd.Flags().StringVarP(
		&notes,
		"notes",
		"n",
		"\n",
		"Text notes. Note: you can add multiple lines by typing \"\\n\" and wrapping the note in double quotes",
	)
	mediaUpdateCmd.Flags().StringVarP(
		&scoreString, "score", "r", "", "The score of the entry. If your score is in emoji, type 1 for 😞, 2 for 😐 and 3 for 😊",
	)
}
