import * as React from 'react';
import { useNavigate } from 'react-router-dom'
import { styled } from '@mui/material/styles';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import MenuIcon from '@mui/icons-material/Menu';
import Container from '@mui/material/Container';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import Tooltip from '@mui/material/Tooltip';
import MenuItem from '@mui/material/MenuItem';
import Badge from '@mui/material/Badge';
import NotificationsIcon from '@mui/icons-material/Notifications';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import MuiDrawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import Divider from '@mui/material/Divider';
import { mainListItems, secondaryListItems } from './listItems';

const pages = ['Dashboard', 'Transactions', 'Accounts', 'Friends'];
const settings = ['Account', 'Preferences', 'Dashboard', 'Logout'];
const drawerWidth = 240;

const Nav = (props) => {
  const [anchorElNav, setAnchorElNav] = React.useState(null);
  const [anchorElUser, setAnchorElUser] = React.useState(null);

  const navigate = useNavigate();

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };
  const handleOpenUserMenu = (event) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  const [open, setOpen] = React.useState(true);
  const toggleDrawer = () => {
    setOpen(!open);
  };

  const Drawer = styled(MuiDrawer, { shouldForwardProp: (prop) => prop !== 'open' })(
    ({ theme, open }) => ({
      '& .MuiDrawer-paper': {
        position: 'relative',
        whiteSpace: 'nowrap',
        width: drawerWidth,
        transition: theme.transitions.create('width', {
          easing: theme.transitions.easing.sharp,
          duration: theme.transitions.duration.enteringScreen,
        }),
        boxSizing: 'border-box',
        ...(!open && {
          overflowX: 'hidden',
          transition: theme.transitions.create('width', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
          }),
          width: theme.spacing(7),
          [theme.breakpoints.up('sm')]: {
            width: theme.spacing(9),
          },
        }),
      },
    }),
  );

  return (
    <>
      <AppBar position="absolute" open={open}>
          <Container maxWidth="xl">
              <Toolbar 
                disableGutters
                sx={{
                  pr: '24px', // keep right padding when drawer closed
                }}  
              >
              <Typography
                  variant="h6"
                  noWrap
                  component="div"
                  sx={{ mr: 2, display: { xs: 'none', md: 'flex' } }}
              >
                  LOGO
              </Typography>

              <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
                  <IconButton
                  size="large"
                  aria-label="account of current user"
                  aria-controls="menu-appbar"
                  aria-haspopup="true"
                  onClick={handleOpenNavMenu}
                  color="inherit"
                  >
                  <MenuIcon />
                  </IconButton>
                  <Menu
                  id="menu-appbar"
                  anchorEl={anchorElNav}
                  anchorOrigin={{
                      vertical: 'bottom',
                      horizontal: 'left',
                  }}
                  keepMounted
                  transformOrigin={{
                      vertical: 'top',
                      horizontal: 'left',
                  }}
                  open={Boolean(anchorElNav)}
                  onClose={handleCloseNavMenu}
                  sx={{
                      display: { xs: 'block', md: 'none' },
                  }}
                  >
                  {pages.map((page) => (
                      <MenuItem 
                      key={page} 
                      onClick={() => {
                          setAnchorElNav(null);
                          navigate(`/${page}`)
                      }}
                      >{page}</MenuItem>
                  ))}
                  </Menu>
              </Box>
              <Typography
                  variant="h6"
                  noWrap
                  component="div"
                  sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}
              >
                  LOGO
              </Typography>
              <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
                  {pages.map((page) => (
                  <Button
                      key={page}
                      onClick={() => {
                      setAnchorElNav(null);
                      navigate(`/${page}`);
                      }}
                      sx={{ my: 2, color: 'white', display: 'block' }}
                  >{page}</Button>
                  ))}
              </Box>

              <Box sx={{ flexGrow: 0 }}>
                  <IconButton 
                  color="inherit"
                  sx={{ pr: 2 }}
                  >
                  <Badge badgeContent={4} color="secondary">
                      <NotificationsIcon />
                  </Badge>
                  </IconButton>
                  <Tooltip title="Open settings">
                  <IconButton 
                      onClick={handleOpenUserMenu} 
                      sx={{ pl: 2 }}>
                      <Avatar alt="Remy Sharp" src="/static/images/avatar/2.jpg" />
                  </IconButton>
                  </Tooltip>
                  <Menu
                  sx={{ mt: '45px' }}
                  id="menu-appbar"
                  anchorEl={anchorElUser}
                  anchorOrigin={{
                      vertical: 'top',
                      horizontal: 'right',
                  }}
                  keepMounted
                  transformOrigin={{
                      vertical: 'top',
                      horizontal: 'right',
                  }}
                  open={Boolean(anchorElUser)}
                  onClose={handleCloseUserMenu}
                  >
                  {settings.map((setting) => (
                      <MenuItem key={setting} onClick={() => {
                          setAnchorElUser(null);
                          navigate(`/${setting}`);
                      }}>
                      <Typography textAlign="center">{setting}</Typography>
                      </MenuItem>
                  ))}
                  </Menu>
              </Box>
              </Toolbar>
          </Container>
      </AppBar>
      {/* <Drawer variant="permanent" open={open}>
        <Toolbar
          sx={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'flex-end',
              px: [1],
          }}
        >
        <IconButton onClick={toggleDrawer}>
            <ChevronLeftIcon />
        </IconButton>
        </Toolbar>
        <Divider />
        <List component="nav">
        {mainListItems}
        <Divider sx={{ my: 1 }} />
        {secondaryListItems}
        </List>
      </Drawer> */} 
    </>
  );
};

export default Nav;