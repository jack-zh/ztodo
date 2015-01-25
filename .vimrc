" jack.zh vim config

set nu
set nocompatible              " be iMproved, required
filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" alternatively, pass a path where Vundle should install plugins
"call vundle#begin('~/some/path/here')

" let Vundle manage Vundle, required
Plugin 'gmarik/Vundle.vim'

" The following are examples of different formats supported.
" Keep Plugin commands between vundle#begin/end.
" plugin on GitHub repo
Plugin 'tpope/vim-fugitive'

Plugin 'L9'

" show file tree on window
Plugin 'scrooloose/nerdtree'

" Git plugin not hosted on GitHub
Plugin 'git://git.wincent.com/command-t.git'

" The sparkup vim script is in a subdirectory of this repo called vim.
" Pass the path to set the runtimepath properly.
Plugin 'rstacruz/sparkup', {'rtp': 'vim/'}

"go-vim
Plugin 'fatih/vim-go'

" YCM
Plugin 'Valloric/YouCompleteMe'

"file tags bar
Plugin 'majutsushi/tagbar'

"
Plugin 'SirVer/ultisnips'


" format with goimports instead of gofmt
let g:go_fmt_command = "goimports"

" disable fmt on save
let g:go_fmt_autosave = 1


" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on    " required
" To ignore plugin indent changes, instead use:
"filetype plugin on
"
" Brief help
" :PluginList       - lists configured plugins
" :PluginInstall    - installs plugins; append `!` to update or just :PluginUpdate
" :PluginSearch foo - searches for foo; append `!` to refresh local cache
" :PluginClean      - confirms removal of unused plugins; append `!` to auto-approve removal
"
" gotags
let g:tagbar_type_go = {
 \ 'ctagstype' : 'go',
 \ 'kinds'     : [
     \ 'p:package',
     \ 'i:imports:1',
     \ 'c:constants',
     \ 'v:variables',
     \ 't:types',
     \ 'n:interfaces',
     \ 'w:fields',
     \ 'e:embedded',
     \ 'm:methods',
     \ 'r:constructor',
     \ 'f:functions'
 \ ],
 \ 'sro' : '.',
 \ 'kind2scope' : {
     \ 't' : 'ctype',
     \ 'n' : 'ntype'
 \ },
 \ 'scope2kind' : {
     \ 'ctype' : 't',
     \ 'ntype' : 'n'
 \ },
 \ 'ctagsbin'  : 'gotags',
 \ 'ctagsargs' : '-sort -silent'
 \ }

" tags and file tree config
nmap <F8> :TagbarToggle<CR>
map <C-n> :NERDTreeToggle<CR>
autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTreeType") && b:NERDTreeType == "primary") | q | endif
autocmd StdinReadPre * let s:std_in=1
autocmd VimEnter * if argc() == 0 && !exists("s:std_in") | NERDTree | endif

" last edit line
au BufReadPost * if line("'\"") > 0|if line("'\"") <= line("$")|exe("norm '\"")|else|exe "norm $"|endif|endif  
