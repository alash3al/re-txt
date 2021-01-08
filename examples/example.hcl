resource "aws_instance" "example" {
    ami = "abc123"

    service {
        key = "value"
    }
}