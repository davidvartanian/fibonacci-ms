Feature: Calculate fibonacci numbers
  As an API user
  I want to be able to send fibonacci positions to an endpoint and receive calculated values
  So that I have access to fibonacci numbers

  Scenario Outline: calculate a single fibonacci number
    When a "GET" request is sent to the endpoint "/get?pos=<pos>"
    Then the HTTP-response code should be "200"
    And the response content should be "<result>"
    Examples:
      | pos | result      |
      | 1   | 1           |
      | 3   | 3           |
      | 5   | 8           |
      | 7   | 21          |
      | 50  | 20365011074 |

  Scenario Outline: calculate a range of fibonacci numbers
    When a "GET" request is sent to the endpoint "/list?min=<min>&max=<max>"
    Then the HTTP-response code should be "200"
    And the response content should be "<result>"
    Examples:
      | min | max | result    |
      | 0   | 3   | [1,1,2,3] |
      | 3   | 6   | [3,5,8,13]|
      | 5   | 10  | [8,13,21,34,55,89]|

  Scenario: deal with invalid request on single fibonacci number
    When a "GET" request is sent to the endpoint "/get"
    Then the HTTP-response code should be "400"
    And the response content should be "Error: missing pos parameter"

  Scenario: deal with invalid request on range of fibonacci numbers
    When a "GET" request is sent to the endpoint "/list"
    Then the HTTP-response code should be "400"
    And the response content should be "Error: missing min or max parameters"