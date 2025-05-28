Feature: user profile photo
  As a user, I want to provide my avatar to make my actions more visible

  Background:
    Given user "Alice" has been created with default attributes


  Scenario Outline: add profile photo
    When user "Alice" sets profile photo to "<file>" using the Graph API
    Then the HTTP status code should be "<http-status-code>"
    And for user "Alice" the profile photo should contain file "<file>"
    Examples:
      | file                                 | http-status-code |
      | filesForUpload/testavatar.jpg        | 200              |
      | filesForUpload/testavatar.png        | 200              |
      | filesForUpload/example.gif           | 200              |
      | filesForUpload/lorem.txt             | 400              |
      | filesForUpload/simple.pdf            | 400              |
      | filesForUpload/broken-image-file.png | 400              |


  Scenario: user tries to get profile photo when none is set
    When user "Alice" tries to get a profile photo using the Graph API
    Then the HTTP status code should be "404"


  Scenario Outline: get profile photo
    Given user "Alice" has set the profile photo to "<file>"
    When user "Alice" gets a profile photo using the Graph API
    Then the HTTP status code should be "200"
    And the profile photo should contain file "<file>"
    Examples:
      | file                          |
      | filesForUpload/testavatar.jpg |
      | filesForUpload/testavatar.png |
      | filesForUpload/example.gif    |


  Scenario Outline: change profile photo
    Given user "Alice" has set the profile photo to "filesForUpload/testavatar.jpg"
    When user "Alice" changes the profile photo to "<file>" using the Graph API
    Then the HTTP status code should be "<http-status-code>"
    And for user "Alice" the profile photo should contain file "<file>"
    Examples:
      | file                                 | http-status-code |
      | filesForUpload/testavatar.jpg        | 200              |
      | filesForUpload/testavatar.png        | 200              |
      | filesForUpload/example.gif           | 200              |
      | filesForUpload/lorem.txt             | 400              |
      | filesForUpload/simple.pdf            | 400              |
      | filesForUpload/broken-image-file.png | 400              |


  Scenario Outline: delete profile photo
    Given user "Alice" has set the profile photo to "<file>"
    When user "Alice" deletes the profile photo using the Graph API
    Then the HTTP status code should be "200"
    When user "Alice" tries to get a profile photo using the Graph API
    Then the HTTP status code should be "404"
    Examples:
      | file                          |
      | filesForUpload/testavatar.jpg |
      | filesForUpload/testavatar.png |
      | filesForUpload/example.gif    |
