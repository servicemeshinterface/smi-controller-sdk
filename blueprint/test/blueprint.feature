Feature: Test Blueprint
  In order to ensure that the Blueprint creates the required resources
  I should test the blueprint

Scenario: Standard Blueprint
    Given I have a running blueprint
    Then the following resources should be running
      | name                | type        |
      | dc1                 | network     |
      | dc1                 | k8s_cluster |
      | consul              | ingress     |
      | grafana             | ingress     |  
      | consul              | helm        |  
      | grafana             | helm        |
      | loki                | helm        |
      | prometheus_stack    | helm        | 
    And a HTTP call to "http://localhost:8500/v1/status/leader" should result in status 200
    And a HTTP call to "http://localhost:8080" should result in status 200
