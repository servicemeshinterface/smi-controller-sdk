Feature: TrafficSplitter
  In order to test the TrafficTarget
  As a developer
  I need to ensure the specification is accepted by the server

  Scenario: Apply alphav4 TrafficTarget
    Given the server is running
    When I create an "alphav4" TrafficSplitter
    Then I expect the controller to have received the details

  Scenario: Apply alphav3 TrafficTarget
    Given the server is running
    When I create an "alphav3" TrafficSplitter
    Then I expect the controller to have received the details
  
  Scenario: Apply alphav2 TrafficTarget
    Given the server is running
    When I create an "alphav2" TrafficSplitter
    Then I expect the controller to have received the details
  
  Scenario: Apply alphav1 TrafficTarget
    Given the server is running
    When I create an "alphav1" TrafficSplitter
    Then I expect the controller to have received the details

  