Feature: TrafficSplitter
  In order to test the TrafficTarget
  As a developer
  I need to ensure the specification is accepted by the server

  Scenario: Apply TrafficTarget
    Given the server is running
    When I create a TrafficSplitter
    Then I expect the controller to have received the details