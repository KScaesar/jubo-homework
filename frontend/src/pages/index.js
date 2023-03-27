import Head from 'next/head'
import {Inter} from 'next/font/google'
import PatientsList from "@/components/PatientsList";
import Paper from "@mui/material/Paper";

const inter = Inter({subsets: ['latin']})

const styles = {
  paper: {
    backgroundColor: 'white',
    height: '100vh',
    padding: '50px',
  },
};

export default function Home() {
  return (
    <>
      <Head>
        <title>jubo-homework-frontend</title>
      </Head>
      <Paper elevation={3} style={styles.paper}>
        <PatientsList/>
      </Paper>
    </>
  );
}