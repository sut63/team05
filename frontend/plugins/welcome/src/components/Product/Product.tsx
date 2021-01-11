import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
  Content,
  Header,
  Page,
  pageTheme /*ContentHeader,*/,
} from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import SaveIcon from '@material-ui/icons/Save';
import { Grid, Container } from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';

import { EntGender } from '../../api/models/EntGender'; // import interface Gender
import { EntGroupOfAge } from '../../api/models/EntGroupOfAge'; // import interface GroupOfAge
import { EntOfficer} from '../../api/models/EntOfficer'; // import interface Officer

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    paper: {
      marginTop: theme.spacing(2),
      marginBottom: theme.spacing(2),
    },
    button: {
      margin: theme.spacing(1),
    },
    textField: {
      width: '25ch',
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    formControl: {
      width: 300,
    },
    
  }),
);

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ระบบประกันสุขภาพ' };
  const api = new DefaultApi();
 
  const [genders, setGenders] = useState<EntGender[]>([]);
  const [groupOfAges, setGroupOfAges] = useState<EntGroupOfAge[]>([]);
  const [officers, setOfficers] = useState<EntOfficer[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [genderID, setGenderID] = React.useState(Number);
  const [groupOfAgeID, setGroupOfAgeID] = React.useState(Number);
  const [officerID, setOfficerID] = React.useState(Number);

  const [productName, setProductName] = React.useState(String);
  const [productPrice, setProductPrice] = React.useState(Number);
  const [productTime, setProductTime] = React.useState(Number);
  const [productPaymentOfYear, setProductPaymentOfYear] = React.useState(Number);

  useEffect(() => {
    const getGender = async () => {
      const gd = await api.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      setGenders(gd);
    };
    getGender();

    const getGroupOfAge = async () => {
      const goa = await api.listGroupofage({ limit: 10, offset: 0 });
      setLoading(false);
      setGroupOfAges(goa);
    };
    getGroupOfAge();

    const getOfficer = async () => {
      const ofc = await api.listOfficer({ limit: 10, offset: 0 });
      setLoading(false);
      setOfficers(ofc);
    };
    getOfficer();
  }, [loading]);


  const CreateProduct = async () => {
    const product = {

      gender: genderID,
      groupOfAge: groupOfAgeID,
      officer: officerID,

      productName: String(productName),
      productPrice: Number(productPrice),
      productTime: Number(productTime),
      productPaymentOfYear: Number(productPaymentOfYear),
    };
    console.log(product);

    const res: any = await api.createProduct({ product: product });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
   
  };

  const GenderhandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setGenderID(event.target.value as number);
  };

  const GroupOfAgehandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setGroupOfAgeID(event.target.value as number);
  };

  const OfficerhandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setOfficerID(event.target.value as number);
  };

  const ProductNamehandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setProductName(event.target.value as string);
  };

  const ProductPricehandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setProductPrice(event.target.value as number);
  };

  const ProductTimehandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setProductTime(event.target.value as number);
  };

  const ProductPaymentOfYearhandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setProductPaymentOfYear(event.target.value as number);
  };
  

  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`${profile.givenName || 'to Backstage'}`}
        subtitle="ระบบสั่งซื้ออุปกรณ์."
      >
       
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              {status ? (
                <div>
                  {alert ? (
                    <Alert severity="success" style={{ width: 305 }}>
                      สั่งซื้อสำเร็จ :)
                    </Alert>
                  ) : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ไม่สำเร็จ !
                    </Alert>
                  )}
                </div>
              ) : null}
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ชื่อผลิตภัณฑ์</div>
            </Grid>
            <Grid item xs={8}>
              <form className={classes.container} noValidate>
                <TextField
                 // label="ป้อนชื่อผลิตภัณฑ์"
                  name="productpament"
                  type=" "
                  value={productName}
                  onChange={ProductNamehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>จำนวนเงินเอาประกัน</div>
            </Grid>
            <Grid item xs={8}>
              <form className={classes.container} noValidate>
                <TextField
                //  label="ป้อนจำนวนเงินเอาประกัน"
                  name="productprie"
                  type=" "
                  value={productPrice}
                  onChange={ProductPricehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ระยะเวลาในการชำระ</div>
            </Grid>
            <Grid item xs={8}>
              <form className={classes.container} noValidate>
                <TextField
                //  label="ป้อนระยะเวลาในารชำระ"
                  name="producttime"
                  type=" "
                  value={productTime}
                  onChange={ProductTimehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ชำระเบี้ยประกัน:ปี</div>
            </Grid>
            <Grid item xs={8}>
              <form className={classes.container} noValidate>
                <TextField
                  name="producttime"
                  type=" "
                  value={productPaymentOfYear}
                  onChange={ProductPaymentOfYearhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>


            <Grid item xs={4}>
              <div className={classes.paper}>เพศ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="gender"
                  value={genderID}
                  onChange={GenderhandleChange}
                >
                  {genders.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.genderName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ช่วงอายุ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="groupofagr"
                  value={groupOfAgeID}
                  onChange={GroupOfAgehandleChange}
                >
                  {groupOfAges.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.groupOfAgeAge}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ผู้บันทึกข้อมูล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="officer"
                  value={officerID}
                  onChange={OfficerhandleChange}
                >
                  {officers.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.officerName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="small"
                className={classes.button}
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreateProduct();
                }}
              >
                Save
              </Button>
              <Button
                size="small"
                style={{ marginLeft: 150 }}
                component={RouterLink}
                to="/result"
                variant="contained"
              >
                Back
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
}