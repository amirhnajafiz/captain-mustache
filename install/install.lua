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
local goBuild = sh.command('go', 'build' , '-o', 'captain-mustache');
local chmodCommand = sh.command('chmod', '+x');

print('Cloning ...');

--
-- clone into repository
--
cloneCommand(GIT_REPO);

print('Building ...');

--- build
goBuild('./captain-mustache/main.go');

local renameCommand = sh.command('mv', './captain-mustache/main', './captain-mustache/captain-mustache');
renameCommand();

chmodCommand('./captain-mustache/captain-mustache');

print(MESSAGE);
print(EXPORT);
