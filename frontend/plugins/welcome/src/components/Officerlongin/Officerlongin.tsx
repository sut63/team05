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

import { EntOfficer} from '../../api/models/EntOfficer'; // import interface Officer


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

  const [officers, setOfficers] = useState<EntOfficer[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);

  const [officeremail, setOfficerEmail] = useState(String);
  const [officerpassword, setOfficerPassword] = useState(String);

  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const getOfficers = async () => {
      const ofc: any = await api.listOfficer({});
      setLoading(false);
      setOfficers(ofc);
    }

    getOfficers();

    const resetOfficerData = async () => {
      setLoading(false);
      localStorage.setItem("officerdata", JSON.stringify(null));
      
    }
    resetOfficerData();

  }, [loading]);

  const EmailthandleChange = (event: any) => {
    setOfficerEmail(event.target.value as string);
  };

  const PasswordthandleChange = (event: any) => {
    setOfficerPassword(event.target.value as string);
  };

  const LoginChecker = async () => {
    officers.map((item: any) => {
      console.log(item.officerEmail);
      if ((item.officerEmail == officeremail) && (item.officerPassword == officerpassword)) {
        setAlert(true);
        localStorage.setItem("officerdata", JSON.stringify(item.id));
        localStorage.setItem("positiondata",JSON.stringify(item.edges.position.positionName));
        history.pushState("", "", "/Product");
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
            value={officeremail}
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
            value={officerpassword}
            onChange={PasswordthandleChange}
            autoComplete="current-password"
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
          </Grid>
        </form>
      </div>
    </Container>
  );
};


