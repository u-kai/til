resource "aws_instance" "sample-ec2" {
    ami = "ami-0e8c2e9fcd76e3b1f"
    instance_type = "t2.micro"
    key_name = "sample-key"
    vpc_security_group_ids = ["${aws_security_group.sample-sg.id}"]
    tags = {
        Name = "sample-ec2"
    }
    user_data = <<EOF
#!/bin/bash
yum install -y httpd
systemctl start httpd
systemctl enable httpd
EOF
}