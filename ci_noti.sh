#!/bin/sh


curl -i -X POST -H 'Content-Type:application/json' -d '{"tag": "text", "text": {"content": "hello '${GITLAB_USER_LOGIN}'\n The pipeline '${CI_PROJECT_NAME}' - '${CI_COMMIT_BRANCH}' 结果是 '${CI_RESULT}' \n Pipeline_url:'${CI_PIPELINE_URL}'\n ci_job_url: '${CI_JOB_URL}'","mentioned_email_list": ["'"${GITLAB_USER_EMAIL}"'"]}}'  ${SEATALK_WEBHOOK_URL}