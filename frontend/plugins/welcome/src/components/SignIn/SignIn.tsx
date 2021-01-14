import React, { useState, useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import ContentHeader from '@material-ui/core';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { DefaultApi } from '../../api/apis';
import { Alert } from '@material-ui/lab';

import { EntMember } from '../../api/models/EntMember';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

export default function Login(props: any) {
  const classes = useStyles();
  const api = new DefaultApi();

  const [members, setMembers] = useState<EntMember[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);

  const [memberemail, setmemberEmail] = useState(String);
  const [memberpassword, setmemberPassword] = useState(String);

  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const getMembers = async () => {
      const res: any = await api.listMember({});
      setLoading(false);
      setMembers(res);
    }

    getMembers();

    const resetMemberData = async () => {
      setLoading(false);
      localStorage.setItem("memberdata", JSON.stringify(null));
      localStorage.setItem("positiondata", JSON.stringify(null));
    }
    resetMemberData();

  }, [loading]);

  const EmailthandleChange = (event: any) => {
    setmemberEmail(event.target.value as string);
  };

  const PasswordthandleChange = (event: any) => {
    setmemberPassword(event.target.value as string);
  };

  const LoginChecker = async () => {
    members.map((item: any) => {
      console.log(item.memberEmail);
      if ((item.memberEmail == memberemail) && (item.memberPassword == memberpassword)) {
        setAlert(true);
        localStorage.setItem("memberdata", JSON.stringify(item.id));
        localStorage.setItem("positiondata", JSON.stringify(item.edges.position.positionName))
        history.pushState("", "", "/Insurance");
        window.location.reload(false);

      }
    })
    setStatus(true);
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      {status ? (
            <div>
              {alert ? (
                <Alert severity="success" onClose={() => { }}>
                  เข้าสู่ระบบสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="error" onClose={() => { setStatus(false) }}>
                    เข้าสู่ระบบไม่สำเร็จ กรุณาตรวจสอบ Email หรือ Password
                  </Alert>
                )}
            </div>
          ) : null}
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            value={memberemail}
            onChange={EmailthandleChange}
            autoComplete="email"
            autoFocus
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            value={memberpassword}
            onChange={PasswordthandleChange}
            autoComplete="current-password"
          />
          <FormControlLabel
            control={<Checkbox value="remember" color="primary" />}
            label="Remember me"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={() => {
              LoginChecker();
            }}
          >
            Sign In

          </Button>
          <Grid container>
            <Grid item xs>
              <Link href="#" variant="body2">
                Forgot password?
              </Link>
            </Grid>
            <Grid item>
              <Link href="#" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
  );
};