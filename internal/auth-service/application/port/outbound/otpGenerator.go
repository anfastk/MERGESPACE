package outbound

type OTPGenerator interface {
    Generate() (string, error)
}
