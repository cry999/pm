import { NextPage } from 'next';
import { AppProps } from 'next/app';
import React, { useEffect } from 'react';

import CssBaseline from '@material-ui/core/CssBaseline';
import {
  StylesProvider as MaterialUIStylesProvider,
  ThemeProvider as MaterialUIThemeProvider,
} from '@material-ui/styles';

import { AuthContextConsumer, AuthContextProvider } from '../src/contexts/auth';
import theme from '../styles/theme';

const App: NextPage<AppProps, {}> = ({ Component, pageProps }) => {
  useEffect(() => {
    const jssStyles = document.querySelector('#jss-server-side');
    if (jssStyles && jssStyles.parentNode) {
      jssStyles.parentNode.removeChild(jssStyles);
    }
  }, []);
  return (
    <MaterialUIStylesProvider injectFirst>
      <MaterialUIThemeProvider theme={theme}>
        <CssBaseline />
        <AuthContextProvider>
          <AuthContextConsumer>
            {(authProps) => <Component {...{ ...pageProps, ...authProps }} />}
          </AuthContextConsumer>
        </AuthContextProvider>
      </MaterialUIThemeProvider>
    </MaterialUIStylesProvider>
  );
};

export default App;
