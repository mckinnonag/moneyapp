import * as React from 'react';
import { createTheme } from '@mui/material/styles';
import { PaletteMode } from '@mui/material';
import { Link as RouterLink, LinkProps as RouterLinkProps } from 'react-router-dom';
import { LinkProps } from '@mui/material/Link';

// Allows the app to use react-router for link elements
const LinkBehavior = React.forwardRef<
  HTMLAnchorElement,
  Omit<RouterLinkProps, 'to'> & { href: RouterLinkProps['to'] }
>((props, ref) => {
  const { href, ...other } = props;
  // Map href (MUI) -> to (react-router)
  return <RouterLink ref={ref} to={href} {...other} />;
});

const getDesignTokens = () => ({
    components: {
      MuiLink: {
        defaultProps: {
          component: LinkBehavior,
        } as LinkProps,
      },
      MuiButtonBase: {
        defaultProps: {
          LinkComponent: LinkBehavior,
        },
      },
    },
    palette: {
        primary: {
        main: '#006064',
        },
        secondary: {
        main: '#ffd54f',
        },
        background: {
        default: '#fafafa',
        paper: '#ffffff',
        },
      },
    });

const Theme = createTheme(getDesignTokens());

export default Theme;