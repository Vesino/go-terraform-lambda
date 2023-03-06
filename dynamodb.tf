resource "aws_dynamodb_table" "profiles" {
  name         = "tutorial_profiles"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "email"

  attribute {
    name = "email"
    type = "S"
  }
}
