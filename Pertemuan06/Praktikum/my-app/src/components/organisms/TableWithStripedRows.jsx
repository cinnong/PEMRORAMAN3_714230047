import { Card } from "@material-tailwind/react";
import { ButtonAtom } from "../atoms/ButtonAtom";
import { TypographyAtom } from "../atoms/TypographyAtom";
import { useMahasiswa } from "../../hooks/useMahasiswa";

const TABLE_HEAD = [
  "NPM",
  "Nama",
  "Prodi",
  "Fakultas",
  "Alamat",
  "Minat",
  "Matakuliah",
];

export function TableWithStripedRows() {
  const { users, loading, error, retry } = useMahasiswa();

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <TypographyAtom variant="h6" color="gray">
          Loading...
        </TypographyAtom>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex flex-col justify-center items-center h-64 space-y-4">
        <TypographyAtom variant="h6" color="red">
          Gagal mengambil data mahasiswa.
        </TypographyAtom>
        <ButtonAtom color="red" onClick={retry}>
          Coba Lagi
        </ButtonAtom>
      </div>
    );
  }

  return (
    <Card className="h-full w-full overflow-auto p-6">
      <div className="flex justify-end p-4">
        <ButtonAtom color="blue">Tambah Data</ButtonAtom>
      </div>

      <table className="w-full min-w-max table-auto text-left">
        <thead>
          <tr>
            {TABLE_HEAD.map((head) => (
              <th
                key={head}
                className="border-b border-blue-gray-100 bg-blue-gray-50 p-4"
              >
                <TypographyAtom
                  variant="small"
                  color="blue-gray"
                  className="font-normal leading-none opacity-70"
                >
                  {head}
                </TypographyAtom>
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.npm} className="even:bg-blue-gray-50/50 align-top">
              <td className="p-4">{user.npm}</td>
              <td className="p-4">{user.nama}</td>
              <td className="p-4">{user.prodi}</td>
              <td className="p-4">{user.fakultas}</td>
              <td className="p-4">
                <ul>
                  <li>Jalan: {user.alamat.jalan}</li>
                  <li>Kelurahan: {user.alamat.kelurahan}</li>
                  <li>Kota: {user.alamat.kota}</li>
                </ul>
              </td>
              <td className="p-4">
                {user.minat.map((m, i) => (
                  <li key={i}>{m}</li>
                ))}
              </td>
              <td className="p-4">
                {user.mata_kuliah.map((mk, i) => (
                  <li key={i}>
                    {mk.nama} ({mk.kode}) - Nilai: {mk.nilai}
                  </li>
                ))}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </Card>
  );
}
