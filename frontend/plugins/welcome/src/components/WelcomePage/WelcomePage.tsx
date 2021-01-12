import React, { FC } from 'react';
import { Typography, Grid,Container} from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
};

export function CardTeam({ name, id, system }: ProfileProps) {
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
          <CardMedia
            component="img"
            alt="นาย สมชาย ใจดี"
            height="140"
            image="url(https://source.unsplash.com/random)"
            title="นาย สมชาย ใจดี"
          />
          <CardContent>
            <Typography gutterBottom variant="h6" component="h6">
              {system}
            </Typography>
            <Typography gutterBottom variant="h6" component="h6">
              {id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const WelcomePage: FC<{}> = () => {
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบประกันสุขภาพ`}></Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Container>
          <Grid container spacing={1}>
            <Grid item xs={12}></Grid>
              <Grid item xs={4}>
                <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
                <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
                
              </Grid>
              <Grid item xs={4}>
                <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
                <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
              </Grid>

              <Grid item xs={4}>
                <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
                <CardTeam name={"นาย สมชาย ใจดี"} id={"B5012345"} system={"ระบบย่อย..."}></CardTeam>
              </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default WelcomePage;