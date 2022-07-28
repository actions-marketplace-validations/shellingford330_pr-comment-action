# pr-comment-action

A GitHub Action which can create a comment on pull request with given owner, repo, pr_number and filepath.


https://user-images.githubusercontent.com/41990509/181576222-cd81b5eb-3466-4791-8fd8-2b89af78fa4c.mp4


## Input

This action supports following inputs.

### owner

The repository owner on pull request.

- _Required_: `yes`
- _Type_: `string`
- _Example_: `shellingford330`

### repo

The repository name on pull request. Either `open` or `closed`. Default: `open`

- _Required_: `yes`
- _Type_: `string`
- _Example_: `pr-comment-action`

### pr_number

The number of pull request.

- _Required_: `yes`
- _Type_: `number`
- _Example_: `1`

### filepath

The filepath containing the content of the comment.

- _Required_: `yes`
- _Type_: `string`
- _Example_: `content.txt`

### template

The template of pull request comment. The contents of the file are assigned to `{{ . }}`.

- _Required_: `no`
- _Type_: `string`
- _Example_: `The contents of the file is {{ . }}.`

## Output

### url

The url of created pull request comment.

## Example

````yaml
name: PR Comment
on:
  issue_comment:
    types: [created]

jobs:
  create:
    if: ${{ github.event.issue.pull_request }}
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Create output file
        run: echo "Create PR comment successfully!" > output.txt
      - name: Comment PR
        uses: shellingford330/pr-comment-action@v0.0.5
        with:
          owner: ${{ github.event.repository.owner.login }}
          repo: ${{ github.event.repository.name }}
          pr_number: ${{ github.event.issue.number }}
          filepath: output.txt
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
````


### Advanced usage


The file content can be embedded using template.

````yaml
name: PR Comment
on:
  issue_comment:
    types: [created]

jobs:
  create:
    if: ${{ github.event.issue.pull_request }}
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Create output file
        run: echo "Create PR comment successfully!" > output.txt
      - name: Comment PR
        uses: shellingford330/pr-comment-action@v0.0.5
        with:
          owner: ${{ github.event.repository.owner.login }}
          repo: ${{ github.event.repository.name }}
          pr_number: ${{ github.event.issue.number }}
          filepath: output.txt
          template: |
            <details>
            <summary>Show Output</summary>

            ```
            {{ . }}
            ```

            </details>
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
````
