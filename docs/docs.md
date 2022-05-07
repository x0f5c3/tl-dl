# toolbox-download

## Usage
> A tool and a library to download the jetbrains-toolbox

toolbox-download

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`toolbox-download completion`|Generate the autocompletion script for the specified shell|
|`toolbox-download help`|Help about any command|
# ... completion
`toolbox-download completion`

## Usage
> Generate the autocompletion script for the specified shell

toolbox-download completion

## Description

```
Generate the autocompletion script for toolbox-download for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`toolbox-download completion bash`|Generate the autocompletion script for bash|
|`toolbox-download completion fish`|Generate the autocompletion script for fish|
|`toolbox-download completion powershell`|Generate the autocompletion script for powershell|
|`toolbox-download completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`toolbox-download completion bash`

## Usage
> Generate the autocompletion script for bash

toolbox-download completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(toolbox-download completion bash)

To load completions for every new session, execute once:

#### Linux:

	toolbox-download completion bash > /etc/bash_completion.d/toolbox-download

#### macOS:

	toolbox-download completion bash > /usr/local/etc/bash_completion.d/toolbox-download

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`toolbox-download completion fish`

## Usage
> Generate the autocompletion script for fish

toolbox-download completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	toolbox-download completion fish | source

To load completions for every new session, execute once:

	toolbox-download completion fish > ~/.config/fish/completions/toolbox-download.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`toolbox-download completion powershell`

## Usage
> Generate the autocompletion script for powershell

toolbox-download completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	toolbox-download completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`toolbox-download completion zsh`

## Usage
> Generate the autocompletion script for zsh

toolbox-download completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	toolbox-download completion zsh > "${fpath[1]}/_toolbox-download"

#### macOS:

	toolbox-download completion zsh > /usr/local/share/zsh/site-functions/_toolbox-download

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... help
`toolbox-download help`

## Usage
> Help about any command

toolbox-download help [command]

## Description

```
Help provides help for any command in the application.
Simply type toolbox-download help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 07 May 2022**
