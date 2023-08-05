--
-- constant variables
--
local GIT_REPO = "https://github.com/amirhnajafiz/captain-mustache.git";
local MESSAGE = "In order to use `captain-mustache` in every place on your system, make sure to" ..
                 "the followings to either `~/.zshrc`, `~/.bash_profile`, or `~/.bashrc`. \n";
local EXPORT = '\t export PATH="<path-to-cloned-repository>:$PATH"\n' ..
               '\t export STRAGO_PATH="<path-to-cloned-repository>"\n'


--
-- importing bash module
--
require "luarocks.loader"
local sh = require('sh');

--
-- build our commands
--
local cloneCommand = sh.command('git', 'clone');
local makeFileCommand = sh.command('make');
local echoCommand = sh.command('echo');

--
-- clone into repository
--
cloneCommand(GIT_REPO);

--- build
makeFileCommand('build');

echoCommand(MESSAGE);
echoCommand(EXPORT);
