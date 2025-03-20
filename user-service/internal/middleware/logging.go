package middleware

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// LoggingInterceptor adalah interceptor untuk mencatat request dan response dari gRPC
func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Catat waktu mulai
		startTime := time.Now()

		// Log request
		reqJSON, err := json.Marshal(req)
		if err != nil {
			log.Printf("Error marshaling request: %v", err)
			reqJSON = []byte("error marshaling request")
		}

		methodName := info.FullMethod
		log.Printf("[%s] Method: %s", startTime.Format(time.RFC3339), methodName)
		log.Printf("Request: %s", string(reqJSON))

		// Proses request
		resp, err := handler(ctx, req)

		// Hitung durasi
		duration := time.Since(startTime)

		// Log response
		if err != nil {
			st, _ := status.FromError(err)
			log.Printf("Error: %s", st.Message())
		} else {
			respJSON, jsonErr := json.Marshal(resp)
			if jsonErr != nil {
				log.Printf("Error marshaling response: %v", jsonErr)
				respJSON = []byte("error marshaling response")
			}
			log.Printf("Response: %s", string(respJSON))
		}

		log.Printf("Duration: %s", duration)
		log.Printf("-----------------------------------")

		return resp, err
	}
}

// PrettyLoggingInterceptor adalah interceptor untuk mencatat request dan response dari gRPC dengan format yang lebih rapi
func PrettyLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Catat waktu mulai
		startTime := time.Now()

		// Log request
		reqJSON, err := json.MarshalIndent(req, "", "  ")
		if err != nil {
			log.Printf("Error marshaling request: %v", err)
			reqJSON = []byte("error marshaling request")
		}

		methodName := info.FullMethod
		log.Printf("\n=== REQUEST ===\n")
		log.Printf("Timestamp: %s", startTime.Format(time.RFC3339))
		log.Printf("Method: %s", methodName)
		log.Printf("Payload: \n%s", string(reqJSON))

		// Proses request
		resp, err := handler(ctx, req)

		// Hitung durasi
		duration := time.Since(startTime)

		// Log response
		log.Printf("\n=== RESPONSE ===\n")
		log.Printf("Duration: %s", duration)

		if err != nil {
			st, _ := status.FromError(err)
			log.Printf("Status: ERROR")
			log.Printf("Error: %s", st.Message())
		} else {
			respJSON, jsonErr := json.MarshalIndent(resp, "", "  ")
			if jsonErr != nil {
				log.Printf("Error marshaling response: %v", jsonErr)
				respJSON = []byte("error marshaling response")
			}
			log.Printf("Status: SUCCESS")
			log.Printf("Payload: \n%s", string(respJSON))
		}

		log.Printf("\n===========================================\n")

		return resp, err
	}
}
