# terraform-provider-npm

The provider used the manage the [npm] membership via [Terraform].

## Resource types

This exposes exactly one resource type: `npm_membership`.
It is used like this:

```
resource "npm_membership" "my_org_developers" {
  for_each = toset([
    "user-one",
    "user-two",
    // ...
  ])

  user = each.value
  org = "my-org"
  role = "developer" // or "admin" or "owner"
}
```

Teams cannot be managed using this provider at this time.

[npm]: https://npmjs.com
[terraform]: https://terraform.io/
