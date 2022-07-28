# pr-comment-action

A GitHub Action which can create a comment with given owner, repo, pr_number and filepath.

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

```yaml

```
