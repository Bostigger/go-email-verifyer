# Go-Email Verifyer

Go-Email Verifyer is a Go-based project that checks and validates email domains for the presence of MX, SPF, and DMARC records. It helps verify the integrity and security settings of a given email domain.
Features

    MX Record Check: Validates the existence of MX (Mail Exchange) records which are essential for email communication.

    SPF Record Check: Verifies the Sender Policy Framework (SPF) records that are used to prevent email spoofing.

    DMARC Record Check: Checks for Domain-based Message Authentication, Reporting & Conformance (DMARC) records that are used to protect a domain from email spoofing and phishing attacks.

Usage

You can provide the domain names as inputs to the program via the console. The results will be printed out in the console in a format that includes information on whether the domain has MX, SPF, and DMARC records.

Here is a sample usage and expected output:


``` go run main.go
domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord
github.com
Domain: github.com
Has MX: true
Has SPF: true
SPF Records: [v=spf1 ip4:192.30.252.0/22 include:_netblocks.google.com include:_netblocks2.google.com include:_netblocks3.google.com include:spf.protection.outlook.com include:mail.zendesk.com include:_spf.salesforce.com include:ser
vers.mcsv.net ip4:166.78.69.169 ip4:166.78.69.170 ip4:166.78.71.131 ip4:167.89.101.2 ip4:167.89.101.192/28 ip4:192.254.112.60 ip4:192.254.112.98/31 ip4:192.254.113.10 ip4:192.254.113.101 ip4:192.254.114.176 ip4:62.253.227.114 ~all]
Has DMARC: true
DMARC Records: [v=DMARC1; p=reject; pct=100; rua=mailto:dmarc@github.com]
```
Function Details


```func CheckDomain(url string)```

This function takes a string argument, which is the domain name to check. It checks the given domain for MX, SPF, and DMARC records, and prints the results to the console.

    Parameters
        url: The domain name to check.

    Returns
        This function does not return any value. It prints the results of the check directly to the console.

Errors

Any errors encountered during the lookup process will be printed to the console with a descriptive message. This includes network errors, invalid domain errors, and any other issues that may occur during the DNS lookup process.
Notes

This program checks the domain records by performing DNS lookups. Ensure that the machine running this program has internet connectivity and the permissions necessary to perform DNS lookups.
Contribution

Contributions are always welcome. Please feel free to fork this project, make your changes following the existing code style, and open a pull request.
