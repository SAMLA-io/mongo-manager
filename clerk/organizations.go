package clerk

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
}

func getUserOrganizations(userId string) (*clerk.OrganizationMembershipList, error) {
	orgMemberships, err := user.ListOrganizationMemberships(context.Background(), userId, &user.ListOrganizationMembershipsParams{})

	if err != nil {
		log.Printf("Error getting organization memberships: %v", err)
	}

	return orgMemberships, nil
}

func GetUserOrganizationId(userId string) (string, error) {
	orgMemberships, err := getUserOrganizations(userId)
	if err != nil {
		return "", err
	}
	if len(orgMemberships.OrganizationMemberships) == 0 {
		return "", errors.New("no organization memberships found")
	}
	return orgMemberships.OrganizationMemberships[0].Organization.ID, nil
}
