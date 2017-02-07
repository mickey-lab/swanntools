package main

import (
	"testing"
	"os"
	"encoding/hex"
	"bytes"
	"fmt"
)

func setUpEnvVars() {
	destEnv, userEnv, passEnv := os.Getenv("AUTH_DEST"), os.Getenv("AUTH_USER"), os.Getenv("AUTH_PASS")
	if destEnv != "" && userEnv != "" && passEnv != "" {
		dest, user, pass = destEnv, userEnv, passEnv
	}
}

func TestLoginValuesCanBeDecoded(t *testing.T) {
	_, err := hex.DecodeString(SuccessfulLoginValues)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to decode successful login values to byte array: ", err.Error())
	}

	_, err = hex.DecodeString(FailedLoginValues)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to decode failed login values to byte array: ", err.Error())
	}
}

func TestIntentMessageCorrectlySetsIntentValue(t *testing.T) {
	expected, _ := hex.DecodeString("00000000000000000000010000000a7b000000292300000000001c010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	res := getIntentMessage("7b")

	if !bytes.Equal(expected, res) {
		t.Error("Intent message not as expected")
	}
}

func TestIntentMessageResponseCorrectlySetsIntentValue(t *testing.T) {
	expected, _ := hex.DecodeString("000000010000000a7b000000292300000000001c010000000100961200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	res := getIntentResponseMessage("7b")

	if !bytes.Equal(expected, res) {
		t.Error("Intent message response not as expected")
	}
}

func TestLoginMessageCorrectlySetsValues(t *testing.T) {
	expected, _ := hex.DecodeString("0000000000000000000001000000197c000000000000000000005461646d696e00000000000000000000000000000000000000000000000000000070617373776400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	res := getLoginMessage("admin", "passwd", "7b")

	if !bytes.Equal(expected, res) {
		t.Error("Login message not as expected")
	}
}