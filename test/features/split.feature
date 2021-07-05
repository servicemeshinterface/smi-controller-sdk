Feature: split.smi-spec.io
  In order to test the TrafficTarget
  As a developer
  I need to ensure the specification is accepted by the server

  Scenario: Apply alpha1 TrafficSplit
    Given the server is running
    When I create the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha1
      kind: TrafficSplit
      metadata:
        name: trafficsplit-sample
      spec:
        service: foo
        backends:
          - service: bar
            weight: 50m
          - service: baz
            weight: 50m
    ```
    Then I expect "UpsertTrafficSplit" to be called 1 time
  
  Scenario: Apply alpha2 TrafficSplit
    Given the server is running
    When I create the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha2
      kind: TrafficSplit
      metadata:
        name: trafficsplit-sample
      spec:
        service: foo
        backends:
          - service: bar
            weight: 50
          - service: baz
            weight: 50
    ```
    Then I expect "UpsertTrafficSplit" to be called 1 time
  
  Scenario: Apply alpha3 TrafficSplit
    Given the server is running
    When I create the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha3
      kind: TrafficSplit
      metadata:
        name: ab-test
      spec:
        service: website
        matches:
        - kind: HTTPRouteGroup
          name: ab-test
        backends:
        - service: website-v1
          weight: 0
        - service: website-v2
          weight: 100
    ```
    And I create the following resource
    ```
      kind: HTTPRouteGroup
      metadata:
        name: ab-test
      matches:
      - name: firefox-users
        headers:
        - user-agent: ".*Firefox.*"
    ```
    Then I expect "UpsertTrafficSplit" to be called 1 time
    Then I expect "HTTPRouteGroup" to be called 1 time
  
  Scenario: Apply alpha4 TrafficSplit
    Given the server is running
    When I create the following resource
    ```
      apiVersion: split.smi-spec.io/v1alpha4
      kind: TrafficSplit
      metadata:
        name: ab-test
      spec:
        service: website
        matches:
        - kind: HTTPRouteGroup
          name: ab-test
        backends:
        - service: website-v1
          weight: 0
        - service: website-v2
          weight: 100
    ```
    And I create the following resource
    ```
      kind: HTTPRouteGroup
      metadata:
        name: ab-test
      matches:
      - name: firefox-users
        headers:
        - user-agent: ".*Firefox.*"
    ```
    Then I expect "UpsertTrafficSplit" to be called 1 time
    Then I expect "HTTPRouteGroup" to be called 1 time