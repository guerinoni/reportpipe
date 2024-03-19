# 1. domains

Date: 2024-03-19

## Status

Under Discussion

## Context

We are exploring the implementation of a feature where companies registered on our platform can offer distinct domains 
to their customers. 
This would enable each customer to access specific functionalities, such as viewing open tickets and creating new ones, 
through personalized URLs like my-provider.com/tickets and my-provider.com/new.

## Decision

The proposed workflow is as follows:
- Upon registration, a company or individual user gains the ability to register new customers, each with a customizable domain.
- Users gain access to a panel or dashboard for managing their customer domains and tickets.
- Companies provide their customers with the designated link to access the platform.
- From the customer's perspective, the provided link is their direct access point to the platform.
- If a customer attempts to register on the platform independently, they are automatically registered under the company's account as a subsidiary entity.

## Consequences

Implementing this approach enhances the platform's flexibility, accommodating diverse companies' needs. 
For instance, if user Z collaborates with providers A and B, they can easily provide to Z their own domain, for handling
tickets respectively, `a-z.com/tickets` and `b-z.com/tickets`. Here Z can ask help to A and B to manage their own needs.

Furthermore, if user Z has its own customers utilizing their services, he can utilize our platform (ReportPipe) from 
main site and register his customer C under his account, so that C can access the platform through `z-c.com/tickets` and
receive support from Z.