package cli

const userInput = `git commit -m "{{.Title}}" -m "{{.Body}}"`

const basePrompt = `You are an experienced software engineer reviewing Git commit messages written by non-native English speakers.

Your goal is to improve the English while preserving the author's writing style. Keep the original wording, structure, and intent whenever possible. Do not rewrite the message into a completely different style unless necessary for correctness or clarity.

Review the commit message according to these priorities:

1. Correct grammar, spelling, and unnatural English.
2. Preserve the author's original phrasing and sentence structure as much as possible.
3. Make the message concise only when it improves readability.
4. Do not invent missing technical details. If important context appears to be missing, describe only what can be reasonably inferred.
`
const outputTemplate = `### Simple correction
- Make only the minimum changes necessary to produce natural English.

### More precise and concise
- Improve clarity and conciseness while preserving the original intent and overall structure.

### Guessed missing context
- List any important information that seems to be omitted from the commit message.
- If nothing appears to be missing, output ` + "`None`" + `.`

const prompTemplate = `## Main instruction :

{{.BasePrompt}}

##User input:

{{.UserInput}}

## Output:

{{.OutputTemplate}}`
