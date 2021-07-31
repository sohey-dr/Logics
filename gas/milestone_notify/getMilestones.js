// GitHub系の処理
function getMilestones(repositories) {
  let milestonesList = [];
  for (const repository of repositories) {
    const response = fetch_pullreq_data_by_graphql(repository);
    const json = response.getContentText();
    const parsedJson = JSON.parse(json);
    let milestones = parsedJson.data.repository.milestones.nodes;

    for (const milestone of milestones) {
      issuesStatus = issuesCount(milestone.issues.nodes);
      milestones[milestones.indexOf(milestone)] = Object.assign(
        milestones[milestones.indexOf(milestone)],
        issuesStatus
      );
      delete milestones[milestones.indexOf(milestone)].issues;
    }

    Array.prototype.push.apply(milestonesList, milestones);
  }

  return milestonesList;
}

function issuesCount(issues) {
  open = 0;
  closed = 0;
  for (const issue of issues) {
    switch (issue.state) {
      case "OPEN":
        open++;
        break;
      case "CLOSED":
        closed++;
        break;
      default:
        console.log("Erorr");
    }
  }

  return { open: open, closed: closed };
}

function fetch_pullreq_data_by_graphql(repository) {
  const graphql_query =
    '{\
    repository(owner: "TechBowl-japan", name: "' +
    repository +
    '") {\
      milestones(states: [OPEN], last: 20) {\
        nodes {\
          title\
          progressPercentage\
          url\
          issues(last: 60) {\
            nodes {\
              state\
            }\
          }\
        }\
      }\
    }\
  }';

  const option = buildRequestOption(graphql_query);
  return UrlFetchApp.fetch("https://api.github.com/graphql", option);
}

function buildRequestOption(graphql) {
  return {
    method: "post",
    contentType: "application/json",
    headers: {
      Authorization: "bearer " + githubAccessToken(),
    },
    payload: JSON.stringify({ query: graphql }),
  };
}

function githubAccessToken() {
  return PropertiesService.getScriptProperties().getProperty(
    "GITHUB_ACCESS_TOKEN"
  );
}
