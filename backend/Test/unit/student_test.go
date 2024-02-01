package unit

import (
	"testing"

	"github.com/W0Ra/SE-lab-final-practice/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_StudentValidator(t *testing.T) {
	g := NewWithT(t)

	t.Run(`StudentID is required`, func(t *testing.T) {
		student := entity.Student{
			StudentID: "", //incorrect
			FirstName: "Wora",
			LastName:  "Poro",
			Email:     "wora@mail.com",
			Phone:     "0123456789",
			Url:       "https://www.link.com",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(("Please insert student id")))
	})

	t.Run(`First name max string 10 characters`, func(t *testing.T) {
		student := entity.Student{
			StudentID: "B1234567",
			FirstName: "Woraphonggg", // incorrect
			LastName:  "Poro",
			Email:     "wora@mail.com",
			Phone:     "0123456789",
			Url:       "https://www.link.com",
			GenderID:  1,
		}
		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(("First name must less than 10 character")))
	})

	t.Run(`Last name min string 4 characters`, func(t *testing.T) {
		student := entity.Student{
			StudentID: "B1234567",
			FirstName: "Woraphong",
			LastName:  "Por", // incorrect
			Email:     "wora@mail.com",
			Phone:     "0123456789",
			Url:       "https://www.link.com",
			GenderID:  1,
		}
		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(("Last name must be at least 4 character")))
	})

	t.Run(`Phone number must be 10 character`, func(t *testing.T) {
		student := entity.Student{
			StudentID: "B1234567",
			FirstName: "Woraphong",
			LastName:  "Poro",
			Email:     "wora@mail.com",
			Phone:     "0123", // incorrect
			Url:       "https://www.link.com",
			GenderID:  1,
		}
		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(("Phone number must be 10 character")))
	})
}
