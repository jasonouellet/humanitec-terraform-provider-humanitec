---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "humanitec_key Resource - terraform-provider-humanitec"
subcategory: ""
description: |-
  A key is used by Humanitec to ensure ensure access to Humanitec hosted drivers.
  The key helps Humanitec operator to establish identity against the Humanitec Driver API
---

# humanitec_key (Resource)

A key is used by Humanitec to ensure ensure access to Humanitec hosted drivers.
The key helps Humanitec operator to establish identity against the Humanitec Driver API

## Example Usage

```terraform
resource "humanitec_key" "example" {
  key = "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAx49tM67P+PDklVOXbbGo\nH3Z5KveonK/bjdiAXsgD8TwGG5w9cP4IRSXKFboHS16Sg4CiZOZdBuJmmfFT7VHK\ni/NIThcD0vuF8UoOQV72Fla+Qb315kWxhlxhVVd6kdqQf4SqVthzzExBMfDyYnLl\n12uFy24XVPGWp9yrFOCrI2pX9/F3aUZh4S1/vDq8pdVBaE302v31aQmMboJqgQVf\nUuvDlFsBPnzvjPjVZhlI/pAP6qfySJ2P6yU6RKCE2HlYtGs499Hvuy+GZTBzd/9+\nsZBqJHwtG2Qwh9vu8PNKUqmAqiSOoOKX4H0xz3Nj4SD/6/qPiCW0e/M2Ws/hXJSv\ntTLud8KNHP6u7aNPYg+V/l6cWcsFr/ZOoMMhqzkEOtKaxCH9c0NqCBv7QxxzF5Md\nt2oHyGrg1QiZd4U2BgWToMbyEaUKJ4G0nFPKYfZh7Udcrt7Vpgpci7jd2W73oWzS\nVhaEyCWgZRnZXXicgT8R55OQdSPXyZcLg57tBP4oursMHGYteSOYSw6nOpc+npW+\nishTpHN52g+z0GLsP7YHZ4oggveKK/7ZNUgBLrJrbhBmPsU/xNqu2jewfC3rEO1X\nbIyD6471lEhdiooy8piRl05vv5uJb3A+vPVvHt6l2koCqKGKOYnfY/okxV7rVD0i\ncOVo7D7KNwPy+CNwZIEDJAcCAwEAAQ==\n-----END PUBLIC KEY-----\n"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `key` (String) The public key that is used for authentication.

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `fingerprint` (String) Hexadecimal representation of the SHA256 hash of the DER representation of the key.
- `id` (String) The ID which refers to a specific key.

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `read` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.

## Import

Import is supported using the following syntax:

```shell
terraform import humanitec_key.example key_id
```
